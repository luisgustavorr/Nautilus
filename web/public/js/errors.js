const quill = new Quill("#editor", {
    modules: {
        toolbar: "#toolbar",
    },
    placeholder: 'Escreva aqui o seu comentário...',
    theme: "snow",
});
let toughtsFromError = []
function loadThouhts(toughts) {
    if (toughts.length == 0) {
        $("#error_comments").html(`
       <div id="nothing_to_see">
        <i class="fa-solid fa-robot"></i>
        <h5>Opa, parece que ninguém deixou um comentário aqui....</h5>
       </div>
        `)
        return
    }

    $("#error_comments").html(`
       ${toughts.map(e => {
        return `  <div class="error_comment_parent">
                <div class="error_comment_perfil_image">
                    ${e.creator_profile_picture == '' ?
                `<img src="https://api.dicebear.com/9.x/thumbs/svg?seed=${e.creator_name}&amp;backgroundColor[]&amp;shapeColor=69d2e7,f1f4dc,f88c49" alt="avatar">`
                :
                `<img src="./images/uploaded/profile_pictures/user_1/picture.png" alt="avatar">`

            }
                </div>
                <div class="error_comment">
                    <div class="error_comment_header">
                        <div class="error_comment_header_text">
                            <h5>${e.creator_name}</h5>
                            <p>${formatDateLongWay(e.created_in)}</p>
                        </div>
                        <div class='actions_parents'>
                        <i class="fa-solid fa-pen-to-square"></i>
                        <i class="fa-solid fa-paperclip" title="Arquivos anexados"></i>
                        <i class="fa-solid fa-quote-right quote_comment" title="Responder comentário" data-error_id="${e.id}"></i>
                        </div>
                    </div>
                    <div class="error_comment_body">
                    ${e.thought}
                    </div>
                </div>

            </div>`
    }).join(`<div class="error_comment_splitter"></div>`)}
        `)
    $("#error_summary_body").html(`
       ${toughts.map(e => {
        return `  <div class="error_summary_row">
                        <div class="error_summary_row_writer_name">
                            <span>${e.creator_name}</span>
                        </div>
                        <div class="error_summary_row_message">
                            <i class="fa-solid fa-turn-up"></i>
                            <div class='message_text'>${e.thought} </div>
                        </div>
                        <div class="error_summary_row_date">
                           ${formatDate(e.created_in)}
                        </div>
                    </div>`
    }).join("")}
        `)

}
if (BindInfos != undefined) {
    if (BindInfos.user_profile_picture == '') {
        $('#add_comment_header .error_comment_perfil_image').html(`<img src="https://api.dicebear.com/9.x/thumbs/svg?seed=${BindInfos.user_name}&amp;backgroundColor[]&amp;shapeColor=69d2e7,f1f4dc,f88c49" alt="avatar">`)
    } else {
        $('#add_comment_header .error_comment_perfil_image').html(`<img src="./images/uploaded/profile_pictures/user_1/picture.png" alt="avatar">`)

    }
}

if (BindInfos.selected_error != undefined) {
    let err = BindInfos.selected_error
    let tags = BindInfos.selected_error.tags
    $('#error_infos_title h3').html(`${err.title} <span>#${err.id}</span>`)
    $('#error_infos_date_creation').text(formatDate(err.created_in))
    $("#error_infos_tags").html(`
        ${tags != undefined && tags != null ? tags.map(e => {
        return `<div class="tag" style="color: ${e.color}; background: ${e.background};">
                      ${e.name}
                    </div>`
    }).join('') : ``}`)
    fetch('/get_error_thoughts/' + err.id, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer '
        }
    })
        .then(response => response.json())
        .then(data => {
            toughtsFromError = data
            loadThouhts(data)
        });
} else {
    console.log('Sem id')

}
$('body').on('click', '.tooltip_opener', function (event) {
    let x = parseFloat(event.pageX) + 10;
    let y = parseFloat(event.pageY) - parseFloat($("header").outerHeight() - 80)
    $("#tooltip_contrato").css("display", "flex")
    moveTooltip(x, y)

})
function moveTooltip(x, y) {
    $("#tooltip_contrato").css("left", x + "px")
    $("#tooltip_contrato").css("top", y + "px")
}
$('body').on('click', '.quote_comment', function () {
  const id = $(this).data('error_id')
  const thg = toughtsFromError.find(e => e.id == id)

 
  const quotedText = $('<div>').html(thg.thought).text()

  quill.setText('')

  let index = 0


  quill.insertText(index, `${thg.creator_name} :`, { bold: true })
  index += `${thg.creator_name} :`.length


  quill.insertText(index, '\n')
  quill.formatLine(index, 1, 'blockquote', true)
  index += 1
  quill.insertText(index, quotedText)
  quill.formatLine(index, quotedText.length, 'blockquote', true)
  index += quotedText.length

 
  quill.insertText(index, '\n')
  quill.formatLine(index, 1, 'blockquote', true)
  index += 1
quill.formatLine(index, index+1, 'blockquote', false)

  quill.setSelection(index, 0, 'user')
})
