window.addEventListener('load', (event) => {
    var linkEl = document.getElementById('full-size-link');
    var imgEl = document.getElementById('full-size');
    document.querySelectorAll('.gallery-id').forEach(
        function(el, i) {
            el.addEventListener('click', (event) => {
                linkEl.href = el.dataset.imageUri;
                imgEl.src = el.dataset.imageUri;
                linkEl.classList.remove('hidden');
                linkEl.scrollIntoView();
            });
        }
    )
})
