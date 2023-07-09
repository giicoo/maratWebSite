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
})