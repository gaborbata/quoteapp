(function() {
    addEventListenerByClass('delete-quote', 'click', function(event) {
        event.preventDefault();
        picoModal({
            content:
                '<div style="text-align:center;">' +
                '<p>Do you really want to delete the quote?</p>' +
                '<button>Yes</button> <button>No</button>' +
                '</div>',
            closeButton: false,
            overlayClose: false
        }).afterCreate(function(modal) {
            var elements = modal.modalElem().getElementsByTagName('button');
            elements[0].addEventListener('click', function() {
                event.target.closest('form').submit();
            });
            elements[1].addEventListener('click', modal.close);
        }).afterClose(function(modal) {
            modal.destroy();
        }).show();
    });

    function addEventListenerByClass(className, event, fn) {
        var list = document.getElementsByClassName(className);
        for (var i = 0; i < list.length; i++) {
            list[i].addEventListener(event, fn);
        }
    }

    // element-closest | CC0-1.0 | github.com/jonathantneal/closest
    if (typeof Element.prototype.matches !== 'function') {
        Element.prototype.matches = Element.prototype.msMatchesSelector || Element.prototype.mozMatchesSelector || Element.prototype.webkitMatchesSelector || function matches(selector) {
            var element = this;
            var elements = (element.document || element.ownerDocument).querySelectorAll(selector);
            var index = 0;
            while (elements[index] && elements[index] !== element) {
                ++index;
            }
            return Boolean(elements[index]);
        };
    }
    if (typeof Element.prototype.closest !== 'function') {
        Element.prototype.closest = function closest(selector) {
            var element = this;
            while (element && element.nodeType === 1) {
                if (element.matches(selector)) {
                    return element;
                }
                element = element.parentNode;
            }
            return null;
        };
    }
})();
