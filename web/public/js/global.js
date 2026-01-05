if (BindInfos != undefined) {
    $('#user_project').html(`${BindInfos.user_name} ${BindInfos.app_name != '' ? `/ <strong> ${BindInfos.app_name} </strong>` : ''}`)
    if (BindInfos.user_profile_picture == '') {
        $('#profile_picture').html(`<img src="https://api.dicebear.com/9.x/thumbs/svg?seed=${BindInfos.user_name}&amp;backgroundColor[]&amp;shapeColor=69d2e7,f1f4dc,f88c49" alt="avatar">`)
    } else {
        $('#profile_picture').html(`<img src="./images/uploaded/profile_pictures/user_1/picture.png" alt="avatar">`)

    }

}
function removerAcentos(str) {
    const tabelaSubstituicao = {
        'á': 'a', 'à': 'a', 'ã': 'a', 'â': 'a',
        'é': 'e', 'è': 'e', 'ê': 'e',
        'í': 'i', 'ì': 'i', 'î': 'i',
        'ó': 'o', 'ò': 'o', 'õ': 'o', 'ô': 'o',
        'ú': 'u', 'ù': 'u', 'û': 'u',
        'Á': 'A', 'À': 'A', 'Ã': 'A', 'Â': 'A',
        'É': 'E', 'È': 'E', 'Ê': 'E',
        'Í': 'I', 'Ì': 'I', 'Î': 'I',
        'Ó': 'O', 'Ò': 'O', 'Õ': 'O', 'Ô': 'O',
        'Ú': 'U', 'Ù': 'U', 'Û': 'U'
    };
    return str.replace(/[áàãâéèêíìîóòõôúùûÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/g, match => tabelaSubstituicao[match]);
}
function formatDate(unformattedDate) {
    let formattedDate = moment(unformattedDate).format('DD/MM/YYYY kk:mm')
    return formattedDate
}
function BooleanButton(button, uniqueClass) {
    if (uniqueClass != "") {
        $('.date_order').each(function () {
            if ($(this).attr("id") != $(button).attr("id")) {
                $(this).removeClass('using')
            }
        });
    }
    if ($('.date_order.using').attr('id') == $(button).attr('id')) {
        if ($(button).data("val") != "1") {
            $(button).data("val", "1")
            $(button).html($(button).html().replace(/fa-arrow-down/g, "fa-arrow-up"))
        } else {
            $(button).data("val", "0")
            $(button).html($(button).html().replace(/fa-arrow-up/g, "fa-arrow-down"))
        }
    }
        $(button).addClass('using')

}