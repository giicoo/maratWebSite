$(document).ready(function(){
    $('.delete').click(function (){
        var dataForm = $("#delete-word").serializeArray();
        words = []
        for (const element of dataForm) {
            el = new Map
            el["word"] = element.name
            words.push(el)
        }
        words = JSON.stringify(words)
        console.log(words)
        $.ajax(
            {
                url:"/delete-word", 
                method: "delete",
                data: words,
                success: function(dt){
                    console.log(dt)
                },
                error: function (err){
                    console.log(err)
                },
                contentType : "application/json"
            }
            ) 
    })
})