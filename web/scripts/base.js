/** закрывать при клике вне */
function openPersonPopup() {
    let popupClasses = $("#person-popup").attr('class');
    if (popupClasses.indexOf('no-display') >= 0) {
        $("#person-popup").removeClass('no-display');
    } else {
        $("#person-popup").addClass('no-display');
    }
}

jQuery(function($){
    $(document).mouseup( function(e) {
        var div = $("#person-popup");
        if ( !div.is(e.target) && div.has(e.target).length === 0 ) {
            $("#person-popup").addClass('no-display');
        }
    });
});

function openPage(page) {
    window.location = './' + page + '.html';
}

function openPopup() {
    $('#add-comment').removeClass('no-display');
}

function closePopup() {
    $('#add-comment').addClass('no-display');
    $('#popup-comment').addClass('no-display');
}

function isAuthorized () {
    let authorized = true;
    if (authorized) {
        $('[authorized]').removeClass('no-display');
        $('[not-authorized]').addClass('no-display');
    } else {
        $('[authorized]').addClass('no-display');
        $('[not-authorized]').removeClass('no-display');
    }
}

function showPassword(element) {
    let show = $(element).attr('class');
    if (show.indexOf('private-off') >= 0) {
        $(element).removeClass('private-off');
    } else {
        $(element).addClass('private-off');
    }
}

window.onload = function() {
    $("#header").load('./components/header.html');
    $("#menu").load('./components/menu.html');
    $("#add-comment").load('./components/add-comment.html');
    $("#footer").load('./components/footer.html');

    setTimeout(function() {
        isAuthorized()
    }, 10);
};
