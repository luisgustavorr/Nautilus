let buttonsBlocked = {}
function blockLoadingButton(element, disable) {
  let id = $(element).attr('id')
  if (disable) {
    buttonsBlocked[id] = $(element).html()
  $(element).html(`<i class="fa-solid fa-spinner fa-spin-pulse"></i>`)
    $(element).prop('disabled',true)
  } else {
    $(element).html(buttonsBlocked[id])
    $(element).prop('disabled',false)
		delete buttonsBlocked[id]
  }
}

function formatDateLongWay(unformattedDate) {
  return moment(unformattedDate).format('DD MMM YYYY hh:mm A')
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
function onElementAdded(containerSelector, elementSelector, callback) {
  const target = $(containerSelector);
  const observer = new MutationObserver(function(mutations) {
    mutations.forEach(function(mutation) {
      if (mutation.addedNodes.length) {
        const elements = $(mutation.addedNodes).find(elementSelector);
        elements.each(function() {
          callback(this);
        });
      }
    });
  });

  observer.observe(target[0], {
    childList: true,
    subtree: true
  });
}
onElementAdded('body', 'format-date', function(element) {
  $(element).text(formatDate($(element).text()))
});
function alertar(title, texto, width = window.outerWidth > 600 ? "500px" : "90%", buttons, type) {

  let buttons_obj = buttons !== undefined ? buttons : {
    OK: {
      btnClass: `btn-${type || 'blue'}`,
      action: function () {
        return
      }
    }
  }
  return $.confirm({
    title: title,
    content: texto,
    boxWidth: width,
    useBootstrap: false,
    type: type || undefined,
    buttons: buttons_obj
  });
}
async function confirm(title, text, button, btnClass = "btn-red", confirmFunction = undefined, argsFuncions = [], width = 500, cancelarButton = undefined) {
  const confirmado = new Promise((resolve, reject) => {
    $.confirm({
      title: title,
      typeAnimated: true,
      content: text,
      boxWidth: width + 'px',
      useBootstrap: false,
      buttons: {
        tryAgain: {
          text: button,
          btnClass: btnClass,
          action: function () {
            if (confirmFunction != undefined) {
              confirmFunction(...argsFuncions)
            }
            resolve(true)
          }
        },
        cancelar: cancelarButton == undefined ? function () {
          resolve(false)

        } : cancelarButton
      }
    });

  })
  return await confirmado.then()
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
function QuickStatusAlert(status,message) {
  let title = status == 500 ? 'Ixi...' : 'Sucesso'
  let color = status == 500 ? 'red' : 'green'
  alertar(title, message, undefined, undefined, color)
}