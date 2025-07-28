window.onload = function () {
    $("#header").load('./components/header.html');
    $("#menu").load('./components/menu.html');
    $("#add-comment").load('./components/add-comment.html');
    $("#footer").load('./components/footer.html');

    setTimeout(function() {
        isAuthorized()
    }, 10);

    $('#auth').on('click', function () {
        $('#registration').removeClass('active');
        $('#registration-data').addClass('no-display');
        $(this).addClass('active');
        $('#auth-data').removeClass('no-display');
    });

    $('#registration').on('click', function () {
        $('#auth').removeClass('active');
        $('#auth-data').addClass('no-display');
        $(this).addClass('active');
        $('#registration-data').removeClass('no-display');
    });
}