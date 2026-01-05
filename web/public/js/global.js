if( BindInfos != undefined ){
    $('#user_project').html(`${BindInfos.user_name} ${BindInfos.app_name !=  ''? `/ <strong> ${BindInfos.app_name} </strong>` : ''  }`)
    if (BindInfos.user_profile_picture ==''){
        $('#profile_picture').html(`<img src="https://api.dicebear.com/9.x/thumbs/svg?seed=${BindInfos.user_name}&amp;backgroundColor[]&amp;shapeColor=69d2e7,f1f4dc,f88c49" alt="avatar">`)
    }else{
        $('#profile_picture').html(`<img src="./images/uploaded/profile_pictures/user_1/picture.png" alt="avatar">`)
        
    }

}
function formatDate(unformattedDate){
let formattedDate = moment(unformattedDate).format('DD/MM/YYYY kk:mm')
return formattedDate
}