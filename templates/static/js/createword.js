$(document).ready(function(){
    $('.create').click(function (){
        var dataForm = $("#create-word").serializeArray();
        console.log(dataForm)
        var data = new Map()
        data["word"] = dataForm[0]["value"]
        data["translate"] = dataForm[1]["value"]
        var formData = JSON.stringify(data)
        console.log(formData)
        $.ajax({
            url: "/add-word",
            method: "post", 
            data: formData,
            success: function(dt){
                console.log(dt)
            },
            error: function(err){
                console.log(err)
            }})
    })
})