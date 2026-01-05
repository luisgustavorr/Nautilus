let openedModal = undefined
function createQuerys(token) {
  let query = 'CASE \n'
  $.get('/categorias/' + token, ret => {
    if (ret == null) {
      alertar('Erro', 'Token Inválido')
      return
    }
    if (ret.length == 0) {
      alertar('Erro', 'Sem categorias cadastradas')
      return
    }

    s_recompra = JSON.parse(ret.s_recompra[0].categorias_recompra) || []
    console.log(s_recompra)
    let prazos = Object.keys(s_recompra)
    prazos.sort((a, b) => {
      return b - a
    })
    for (let e of prazos) {
      console.log(e)
      if ((e).toString() == '0') {
             query +=`ELSE 0 \n`
        continue
      }
      query += `WHEN _id_categoria_ IN (${s_recompra[e].join(',')}) THEN ${e} \n`
    }
     query += `END AS duracao`
    console.log(query)
  })

}
async function buildModal(obj, values) {
  if(openedModal != undefined){
    openedModal.close()
  }
  console.log(values)
  function formToJSON(form) {
    let jsonObject = {}
    let inputs = $(form).find("[auto]")
    inputs.each(function () {
      if ($(this).attr("id") != "") {
        if ($(this).attr('class') != undefined && $(this).attr('class').includes('money_mask')) {
          jsonObject[$(this).attr("id")] = $(this).val().toString().replace(/\./g, '').replace(/,/g, '.')
        } else if ($(this).attr('class') != undefined && $(this).attr('class').includes('date_input')) {
          jsonObject[$(this).attr("id")] = moment($(this).val(), "DD/MM/YYYY").format()
        } else if ($(this).attr('type') != undefined && $(this).attr('type').includes('checkbox')) {
          jsonObject[$(this).attr("id")] = $(this).is(':checked')
        } else {
          jsonObject[$(this).attr("id")] = $(this).val()

        }
      }
    })
    return jsonObject
  }
  function verificarForm(elementos, customFunc = undefined) {
    let prosseguir = true
    elementos.each(function () {
      if (!prosseguir) return
      $(this).css('border', '1px solid gray')
      if ($(this).val() == '') {
        $(this).css('border', '1px solid red')
        $(this).focus()
        alertar('Opa', 'Prencha todos os campos obrigatórios')
        if (customFunc != undefined) {
          customFunc($(this))
        }
        prosseguir = false
      }
    })
    return prosseguir
  }
  // cálculo width input_father = ((({quantidade_max_inputs}/{quantidade_inputs})/{quantidade_max_inputs})*100) - ({quantidade_inputs}-1)/{porcentagem_gap}
  let rowsQt = obj.rows.length
  let formHeight = rowsQt * 3 + ((rowsQt - 1) * 0.5)

  formHeight = formHeight > 25 ? 25 : formHeight
  const answerPromise = new Promise((resolve, reject) => {
    openedModal = $.confirm({
      title: obj.title,
      typeAnimated: true,
      content: `
  			<div id="built-form" style="height:${formHeight}rem;">
        	${obj.rows.map(e => {
        let inputs = Object.keys(e)
        return `<div class="built-form-rows">${inputs.map(i => {
          let row = e[i]
          let input = row.input
          let inputInfos = Object.keys(input)
          return `
        			<div class="input_father_${input["type"]}" style="width:${(((4 / inputs.length) / 4) * 100) - (inputs.length - 1) / 2}%;">
              <label for="${i}">${row.label}</label>
              <${row.element} id="${i}" auto="true" ${inputInfos.map(ii => {
            return `${ii}="${input[ii]}"`
          }).join(" ")}>${row.htmlChild || ''}</${row.element}>
              </div>
            `
        }).join('')}</div>`

      }).join('')}
        </div>
      `,
      boxWidth: '60%',
      useBootstrap: false,
      buttons: {
        tryAgain: {
          text: "Salvar",
          btnClass: "btn-green",
          action: function () {
            let formCompleto = verificarForm($('#built-form [necessary="true"]'))
            if (!formCompleto) {
              return false
            }
            resolve(formToJSON($('#built-form')))
            console.log('ENVIADO')
          },
        },
        cancelar: function () {
          resolve(undefined)
          return true
        }
      }, onContentReady: function () {
        // if (onBuiltFunction != undefined) {
        //   onBuiltFunction(...onBuiltFunctionArgs)
        // }
$($('#built-form input')[0]).focus()
        if(values != undefined){
          fillInputsByJSON($('#built-form'),values)
        }
      }
    });
  })
  const answer = await answerPromise
  return {
    answer,
    canceled: answer == undefined
  }
}

function alertar(title, texto, width = window.outerWidth > 600 ? "500px" : "90%", buttons, type) {
  let buttons_obj = buttons !== undefined ? buttons : {
    OK: {
      btnClass: `btn-${type || 'blue'}`,
      action: function () {
        return
      }
    }
  }
  console.log(buttons_obj)
  return $.confirm({
    title: title,
    content: texto,
    boxWidth: width,
    useBootstrap: false,
    type: type || undefined,
    buttons: buttons_obj
  });
}
