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
})();
