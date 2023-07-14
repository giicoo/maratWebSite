$(document).ready(function(){
    $("#btmSumbit").click(function(){
        var formData = JSON.stringify({"login":$('#val1').val(), "password":$('#val2').val()});
        $.ajax(
            {
                url: '/singin',
                method: 'post',
                data: formData,
                success: function(dt){
                    window.location.replace('/')
                },
                error: function (err){
                    console.log(err)
                },
                contentType : "application/json"
            }

        )

    });
    $("#rbtmSumbit").click(function(){
        var formData = JSON.stringify({"login":$('#rval1').val(), "password":$('#rval3').val()});
        $.ajax(
            {
                url: '/singup',
                method: 'post',
                data: formData,
                success: function(dt){
                    window.location.replace('/sing')
                },
                error: function (err){
                    $('.error').css('display', 'block')
                    console.log(err)
                },
                contentType : "application/json"
            }

        )

    });
    $('.change').click(function (){
        if ($('.log').css('display') == 'none'){
            $('.log').css('display', 'block')
            $('.reg').css('display', 'none')
        } else {
            $('.log').css('display', 'none')
            $('.reg').css('display', 'block')
        }
        }
    )
});