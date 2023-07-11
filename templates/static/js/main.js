$(document).ready(function(){
    $('.onward').click(function (){
        var tests = document.getElementsByClassName('word')
        var now = $('meta[name=now]').attr('content');
        if (now != tests.length){
        if (now != 0){
            var last = tests[now-1]
            last.className = "word invisible"
        }
        var test = tests[now]
        

        test.className = "word visible"
        
        $('meta[name=now]').attr('content', parseInt(now)+1);
        }
    }
    )
    $('.submit').click(function (){
        var tests = document.getElementsByClassName('word')
        var now = $('meta[name=now]').attr('content');
        var test = tests[now-1]
        test.className = "word invisible"
    
        var dataForm = $("#form_test").serializeArray();

        
        var data = []
        for (const element of dataForm) {
            console.log(element)
            var el = new Map()
            el.set("word", element.name)
            el.set("translate", element.value)
            data.push(Object.fromEntries(el))
        }
    

        var formData = JSON.stringify(data)
        var test_name = $('meta[name=test_name]').attr('content');
        $.ajax(
            {
                url: '/check-test/'+test_name,
                method: 'post',
                data: formData,
                success: function(dt){
                    var formData2 = JSON.stringify(dt)
                    $.ajax (
                        {
                            url: '/test/res-page/'+test_name,
                            method: 'post',
                            data: formData2,
                            success: function(dt) {
                                document.getElementsByClassName("main")[0].innerHTML = dt
                                
                                console.log(dt)
                            },
                            error: function(err) {
                                console.log(err)
                            }
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
});
