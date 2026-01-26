let lastErrorsRecovered = []
let tags = []

fetch('/get_error_tags/1', {
  method: 'GET',
  headers: {
    'Authorization': 'Bearer '
  }
})
  .then(response => response.json())
  .then(data => {
    tags = data
    $('#select_tag_filter').html(`
                            <option value=""selected> TAG</option>
                   ${tags.map(f => {
      return ` <option value="${f.id}">${f.name}</option>`
    }).join("")}
                `)
  });
$('#last_edit,#date_creation').data('val', "1")
// $.ajax({
//   url: '/get_errors/1',
//   method: 'GET',
//   contentType: 'text/html',
//   success: function (ret) {
//     console.log(ret)
//     $('.tabela_body').html(ret)

//   },
//   error: function (err) {
//     console.log(err)
//   }
// });
function loadErrorsInTable(data) {
  // $('.tabela_body').html("")
  $('.tabela_counter').html(`<i class="fa-solid fa-spinner fa-spin-pulse"></i>`)

  for (const e of data) {
    $('.tabela_body').append(`<div class="error_row">
                <div class="status_column">
                  ${e.verified ? `  <div  class="status_icon_success">
             <i class="fa-solid fa-circle-check"></i>
                    </div>` : `  <div  class="status_icon_error">
<i class="fa-solid fa-clock"></i>
                    </div>`}
                </div>
                <div class="main_column">
                    <a class="error_title" href="/error_${e.id}">
                       ${e.title}
                    </a>
                    <p class="error_description">
                       Criado por : ${e.creator_name} às ${formatDate(e.created_in)}
                    </p>
                    <div class="error_tags">
                    ${e.tags ? e.tags.map(f => {
      return ` <div class="tag" style="color: ${f.color};background: ${f.background};" title="${f.description}">
                            ${f.name}
                        </div>`
    }).join('') : ``}
                       
                       
                    </div>
                </div>
                <div class="infos_column">
                    <div class="infos_column_first_row">
                        <span class="error_id">#${e.id}</span>
                        <div class="date_infos">
                            <span>
                               <i class="fa-solid fa-clock" title = "O erro ocorreu nesse horário"></i> ${formatDate(e.error_occurred_in)}

                            </span>
                            <span>
                                <i class="fa-solid fa-arrows-spin" title = "Última edição no erro"></i> ${formatDate(e.last_edited_in)}

                            </span>
                        </div>
                    </div>
                    <div class="infos_columns_files_row" id_error="${e.id}" >
                      <i class="fa-solid fa-images"></i> ${e.files?.length || 0} arquivo(s) anexado(s)
                    </div>
                </div>
            </div>`)
  }
  $('.tabela_counter').html(data.length == 0 ? '#' : data.length)

}
$('#search_creator_input').on("keyup", function (e) {
  if (e.keyCode == 13) {
    filterErrors(lastErrorsRecovered)
  }
});
function filterErrors(errors) {
  console.log('Filtrando')
  if ($('#select_status_filter').val() != "") {
    console.log('Por status')
    let verificado = $('#select_status_filter').val() == "1"
    if (verificado) {
      $('#select_status_filter').css({
        "background": "green"
      })
    } else {
      $('#select_status_filter').css({
        "background": "red"
      })
    }
    errors = errors.filter(e => {
      return e.verified == verificado
    })

  } else {
    $('#select_status_filter').removeAttr('style')
  }
  if ($('#select_tag_filter').val() != "") {
    let selectedTag = tags.filter(e => e.id == $('#select_tag_filter').val())[0]
    $('#select_tag_filter').css({
      "color": selectedTag.color,
      "background": selectedTag.background
    })
    errors = errors.filter(e => {
      if (e.tags == undefined) return false
      return e.tags.map(f => f.id).includes(parseInt($('#select_tag_filter').val()))
    })

  } else {
    $('#select_tag_filter').removeAttr('style')
  }
  if ($('#search_creator_input').val() != "") {
    console.log('Por nome')
    let termoDeBusca = $('#search_creator_input').val()
    errors = errors.filter(obj => removerAcentos(obj["creator_name"]).toLowerCase().includes(removerAcentos(termoDeBusca).toLowerCase()));
  }
  let multiplier = parseInt($('.date_order.using').data('val')) == 1 ? 1 : -1
  if ($('.date_order.using').attr('id') == 'last_edit') {
    errors.sort((a, b) => {
      if (moment(moment(a.last_edited_in).format()).diff(moment(moment(b.last_edited_in).format()), 'hours') > 0) {
        return 1 * multiplier
      } else if (moment(moment(a.last_edited_in).format()).diff(moment(moment(b.last_edited_in).format()), 'hours') < 0) {
        return -1 * multiplier
      } else {
        return 0
      }
    })
  } else {
    errors.sort((a, b) => {
      if (moment(moment(a.created_in).format()).diff(moment(moment(b.created_in).format()), 'hours') > 0) {
        return 1 * multiplier
      } else if (moment(moment(a.created_in).format()).diff(moment(moment(b.created_in).format()), 'hours') < 0) {
        return -1 * multiplier
      } else {
        return 0
      }
    })
  }
  loadErrorsInTable(errors)

}
// filterErrors(lastErrorsRecovered)
$('#last_edit,#date_creation').click(function () {
  filterErrors(lastErrorsRecovered)
})
$('body').on('click', '.infos_columns_files_row', function () {
  let id = $(this).data('id_error')
  $.ajax({
    url: '/get_files_from_error/' + id,
    method: 'GET',
    contentType: 'application/json',

    success: function (ret) {
      console.log(ret)
      new browse(ret, `/images/uploaded/error_files/error_${id}/`)

    },
    error: function (err) {
      console.log(err)
    }
  });
})