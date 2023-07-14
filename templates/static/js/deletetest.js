$(document).ready(function(){
    $('.delete').click(function (){
        var dataForm = $("#delete-test").serializeArray();
        tests = []
        for (const element of dataForm) {
            el = new Map
            el["name"] = element.name
            tests.push(el)
        }
        tests = JSON.stringify(tests)
        $.ajax(
            {
                url:"/delete-test", 
                method: "delete",
                data: tests,
                success: function(dt){
                    document.location.reload()
                },
                error: function (err){
                    console.log(err)
                },
                contentType : "application/json"
            }
            ) 
    })
})