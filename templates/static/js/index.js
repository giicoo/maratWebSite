$(document).ready(function(){
    $('.logout').click(function (){
        $.ajax({
            url:"/logout",
            method: "post",
            success: function(dt){
                document.location.href = "/"
                console.log(dt)
            },
            error: function(err){
                console.log(err)
            }
        })
    })
    $('.sing-in-up').click(function (){
        document.location.href = "/sing"
    })

    $(".stat").click(function(){
        $.ajax({
            url:"/statics",
            method: "post",
            xhrFields: {
                responseType: 'blob'
            },
            success: function(dt){
                var a = document.createElement('a');
                var url = window.URL.createObjectURL(dt);
                a.href = url;
                a.download = 'statistics.xlsx';
                document.body.append(a);
                a.click();
                a.remove();
                window.URL.revokeObjectURL(url);
            },
            error: function(err){
                console.log(err)
            },
        })
    })
})