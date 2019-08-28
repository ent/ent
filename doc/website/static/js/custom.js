/* eslint-disable */
window.addEventListener('load', function() {
    // add an id tag for images based on their alt attribute.
    document.querySelectorAll('.container img').forEach(function(el) {
        el.id = el.alt;
    });
});
