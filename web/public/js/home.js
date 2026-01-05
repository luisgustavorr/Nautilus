$(document).ready(function () {
    fetch('/get_errors/1', {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer '
        }
    })
        .then(response => response.json())
        .then(data => {
            for (const e of data) {
                $('.tabela_body').append(`    <div class="error_row">
                <div class="status_column">
                  ${e.verified ? `  <div  class="status_icon_success">
             <i class="fa-solid fa-circle-check"></i>
                    </div>` : `  <div  class="status_icon_error">
<i class="fa-solid fa-square-xmark"></i>
                    </div>`}
                </div>
                <div class="main_column">
                    <span class="error_title">
                       ${e.title}
                    </span>
                    <p class="error_description">
                        ${e.message}
                    </p>
                    <div class="error_tags">
                        <div class="tag" style="color: white;background: blue;">
                            teste texto muito maior q o outro
                        </div>
                        <div class="tag" style="color: white;background: orange;">
                            I.A
                        </div>
                        <div class="tag" style="color: white;background: red;">
                            ERROR
                        </div>
                    </div>
                </div>
                <div class="infos_column">
                    <div class="infos_column_first_row">
                        <span class="error_id">#${e.id}</span>
                        <div class="date_infos">
                            <span>
                               <i class="fa-solid fa-clock"></i> ${formatDate(e.created_in)}

                            </span>
                            <span>
                                <i class="fa-solid fa-arrows-spin"></i> ${formatDate(e.last_edited_in)}

                            </span>
                        </div>
                    </div>
                    <div class="infos_columns_files_row">
                        <i class="fa-solid fa-images"></i> ${e.files_count ||0} arquivo(s) anexado(s)
                    </div>
                </div>


            </div>`)
            }
        });
})