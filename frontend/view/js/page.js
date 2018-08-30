
function prevPage(query, currentPage) {

    currentPage--

    var url = "search?q=" + query + "&current=" + currentPage;

    $.ajax({
        url: url,
        type: 'get',
        success: function (data) {
            return true;
        }
    });

    return false;
}


function nextPage(query, currentPage) {

    currentPage++

    var url = "search?q=" + query + "&current=" + currentPage;

    $.ajax({
        url: url,
        type: 'get',
        success: function (data) {
            return true;
        }
    });

    return false;
}
