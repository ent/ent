
/**
 * Copyright 2019-present Facebook Inc. All rights reserved.
 *
 * This source code is licensed under the Apache 2.0 license found
 * in the LICENSE file in the root directory of this source tree.
 *
 * @format
 */

/* eslint-disable */
window.addEventListener('load', function() {
    // add an id tag for images based on their alt attribute.
    document.querySelectorAll('.container img').forEach(function(el) {
        el.id = el.alt;
    });
});
