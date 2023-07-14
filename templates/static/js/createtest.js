$(document).ready(function(){
    $('.create').click(function (){
        var dataForm = $("#create-test").serializeArray();
        var data = new Map()
        data["name"] = dataForm[0]["value"]
        data["words"] = []
        dataForm = dataForm.slice(1)
        words = []
        for (const element of dataForm) {
            el = new Map
            el["word"] = element.name
            words.push(el)
        }
        words = JSON.stringify(words)
        $.ajax(
            {
                url:"/get-words-by-names", 
                method: "post",
                data: words,
                success: function(dt){
                    data["words"] = dt
                    var formData = JSON.stringify(data)
                    $.ajax(
                        {
                            url:"/add-test", 
                            method: "post",
                            data: formData,
                            success: function(dt){
                                document.location.reload()
                            },
                            error: function (err){
                                console.log(err)
                            },
                            contentType : "application/json"
                        }
                    )
                },
                error: function (err){
                    console.log(err)
                },
                contentType : "application/json"
            }
            ) 
    })
})