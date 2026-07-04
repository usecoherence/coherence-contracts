use anyhow::{bail, Context, Result};
use clap::{Parser, Subcommand};
use serde::Serialize;
use std::fs;
use std::path::{Path, PathBuf};
use walkdir::WalkDir;

#[derive(Debug, Parser)]
#[command(name = "coherence-contracts-tools")]
#[command(about = "Small checks for coherence protobuf contracts")]
struct Cli {
    #[command(subcommand)]
    command: Command,
}

#[derive(Debug, Subcommand)]
enum Command {
    Check {
        #[arg(long, default_value = "proto")]
        proto_dir: PathBuf,
    },
    Inventory {
        #[arg(long, default_value = "proto")]
        proto_dir: PathBuf,
    },
    Graph {
        #[arg(long, default_value = "proto")]
        proto_dir: PathBuf,
    },
}

#[derive(Debug, Serialize)]
struct ProtoInventoryItem {
    path: String,
    package: Option<String>,
    imports: Vec<String>,
    messages: Vec<String>,
    enums: Vec<String>,
}

fn main() -> Result<()> {
    let cli = Cli::parse();

    match cli.command {
        Command::Check { proto_dir } => check(&proto_dir),
        Command::Inventory { proto_dir } => inventory(&proto_dir),
        Command::Graph { proto_dir } => graph(&proto_dir),
    }
}

fn check(proto_dir: &Path) -> Result<()> {
    let files = proto_files(proto_dir)?;
    let mut errors = Vec::new();

    for file in files {
        let source = fs::read_to_string(&file)
            .with_context(|| format!("failed to read {}", file.display()))?;

        if !source.contains("syntax = \"proto3\";") {
            errors.push(format!("{}: missing proto3 syntax", file.display()));
        }

        let package = parse_package(&source);
        let expected = expected_package(proto_dir, &file);

        if package.as_deref() != Some(expected.as_str()) {
            errors.push(format!(
                "{}: package {:?}, expected {:?}",
                file.display(),
                package,
                expected
            ));
        }

        if is_kernel_file(&file) {
            for import in parse_imports(&source) {
                if import.contains("coherence/core/") {
                    errors.push(format!(
                        "{}: kernel must not import core file {:?}",
                        file.display(),
                        import
                    ));
                }
            }
        }
    }

    if errors.is_empty() {
        println!("contract checks passed");
        Ok(())
    } else {
        for error in &errors {
            eprintln!("ERROR: {error}");
        }
        bail!("{} contract check(s) failed", errors.len())
    }
}

fn inventory(proto_dir: &Path) -> Result<()> {
    let files = proto_files(proto_dir)?;
    let mut items = Vec::new();

    for file in files {
        let source = fs::read_to_string(&file)?;
        items.push(ProtoInventoryItem {
            path: file.display().to_string(),
            package: parse_package(&source),
            imports: parse_imports(&source),
            messages: parse_declarations(&source, "message"),
            enums: parse_declarations(&source, "enum"),
        });
    }

    println!("{}", serde_json::to_string_pretty(&items)?);
    Ok(())
}

fn graph(proto_dir: &Path) -> Result<()> {
    let files = proto_files(proto_dir)?;
    println!("digraph imports {{");
    println!("  rankdir=LR;");

    for file in files {
        let source = fs::read_to_string(&file)?;
        let from = normalize_proto_path(proto_dir, &file);
        println!("  \"{}\";", from);
        for import in parse_imports(&source) {
            println!("  \"{}\" -> \"{}\";", from, import);
        }
    }

    println!("}}");
    Ok(())
}

fn proto_files(proto_dir: &Path) -> Result<Vec<PathBuf>> {
    let mut files = Vec::new();
    for entry in WalkDir::new(proto_dir) {
        let entry = entry?;
        if entry.file_type().is_file() && entry.path().extension().is_some_and(|ext| ext == "proto")
        {
            files.push(entry.path().to_path_buf());
        }
    }
    files.sort();
    Ok(files)
}

fn parse_package(source: &str) -> Option<String> {
    source.lines().find_map(|line| {
        let trimmed = line.trim();
        trimmed
            .strip_prefix("package ")
            .and_then(|rest| rest.strip_suffix(';'))
            .map(|value| value.trim().to_string())
    })
}

fn parse_imports(source: &str) -> Vec<String> {
    source
        .lines()
        .filter_map(|line| {
            let trimmed = line.trim();
            trimmed
                .strip_prefix("import ")
                .and_then(|rest| rest.trim().strip_prefix('"'))
                .and_then(|rest| rest.split('"').next())
                .map(str::to_string)
        })
        .collect()
}

fn parse_declarations(source: &str, keyword: &str) -> Vec<String> {
    source
        .lines()
        .filter_map(|line| {
            let trimmed = line.trim();
            let rest = trimmed.strip_prefix(keyword)?;
            let name = rest.trim().split([' ', '{']).next()?;
            if name.is_empty() {
                None
            } else {
                Some(name.to_string())
            }
        })
        .collect()
}

fn expected_package(proto_dir: &Path, file: &Path) -> String {
    let rel = file.strip_prefix(proto_dir).unwrap_or(file);
    let parent = rel.parent().unwrap_or_else(|| Path::new(""));
    parent
        .components()
        .map(|component| component.as_os_str().to_string_lossy().replace('-', "_"))
        .collect::<Vec<_>>()
        .join(".")
}

fn normalize_proto_path(proto_dir: &Path, file: &Path) -> String {
    file.strip_prefix(proto_dir)
        .unwrap_or(file)
        .display()
        .to_string()
}

fn is_kernel_file(path: &Path) -> bool {
    let parts = path
        .components()
        .map(|component| component.as_os_str().to_string_lossy().to_string())
        .collect::<Vec<_>>();

    parts
        .windows(3)
        .any(|window| window[0] == "coherence" && window[1] == "kernel" && window[2] == "v1")
}
