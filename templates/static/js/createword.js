$(document).ready(function(){
    $('.create').click(function (){
        var dataForm = $("#create-word").serializeArray();
        var data = new Map()
        data["word"] = dataForm[0]["value"]
        data["translate"] = dataForm[1]["value"]
        var formData = JSON.stringify(data)
        $.ajax({
            url: "/add-word",
            method: "post", 
            data: formData,
            success: function(dt){
                document.location.reload()
            },
            error: function(err){
                console.log(err)
            }})
    })
})