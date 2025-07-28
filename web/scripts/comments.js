function updateComment() {
    console.log('здесь нужно получить полную информацию с бэка и впихнуть в попап. Удачи');
    openPopup();
}
function showAll() {
    $('#popup-comment').removeClass('no-display');
}
function updateSort() {
    let sort = $('#sort').attr('class');
    if (sort.indexOf('up') >= 0) {
        $('#sort').addClass('down');
        $('#sort').removeClass('up');
    } else {
        $('#sort').addClass('up');
        $('#sort').removeClass('down');
    }
}