(() => {
  document.querySelectorAll('.persist-demo').forEach((root) => {
    const ctxRadios = root.querySelectorAll('input[name="ctx"]');
    const flag = root.querySelector('.persist-flag');
    const hint = root.querySelector('.persist-hint');
    if (!ctxRadios.length || !flag || !hint) return;

    const render = () => {
      const ctx = root.querySelector('input[name="ctx"]:checked').value;
      const force = flag.checked;
      const activeCol = force ? 'force' : ctx;

      root.querySelectorAll('[data-col]').forEach((cell) => {
        cell.classList.toggle('active-col', cell.dataset.col === activeCol);
      });

      if (force) {
        hint.textContent = root.dataset.hintForce;
      } else if (ctx === 'ci') {
        hint.textContent = root.dataset.hintCi;
      } else {
        hint.textContent = root.dataset.hintLocal;
      }
    };

    ctxRadios.forEach((radio) => radio.addEventListener('change', render));
    flag.addEventListener('change', render);
    render();
  });
})();
