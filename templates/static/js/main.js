$(document).ready(function(){
    $('.onward').click(function (){
        var tests = document.getElementsByClassName('word_test')
        var now = $('meta[name=now]').attr('content');
        if (now != tests.length){
        if (now != 0){
            var last = tests[now-1]
            last.className = "word_test invisible"
        }
        var test = tests[now]
        

        test.className = "word_test visible"
        
        $('meta[name=now]').attr('content', parseInt(now)+1);
        }
    }
    )
    $('.submit').click(function (){
        var tests = document.getElementsByClassName('word_test')
        var now = $('meta[name=now]').attr('content');
        var test = tests[now-1]
        test.className = "word_test invisible"
    
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
        
        $.ajax(
            {
                url: '/check-test',
                method: 'post',
                data: formData,
                success: function(dt){
                    var res = document.getElementById("result")
                    var x = 0
                    for (let i = 0; i < dt.length; i++) {
                        if (dt[i]["check"]) {
                            var chil = document.createElement("div")
                            chil.setAttribute("class", "right");
                            chil.textContent = dt[i]["word"]["word"] + " - " + dt[i]["right"]
                            res.children[0].appendChild(chil)
                            x ++
                        } else {
                           
                            var chil = document.createElement("div")
                            chil.setAttribute("class", "error");
                            chil.textContent = dt[i]["word"]["word"] + " - " + dt[i]["right"]
                            res.children[0].appendChild(chil)
                        }
                      }
                    res.className = "result visible"
                    proc = document.getElementById("procent")
                    proc.textContent = x/(dt.length)*100 + "%"

                    test = document.getElementById("test")
                    test.className = "test testres"
                },
                error: function (err){
                    console.log(err)
                },
                contentType : "application/json"
            }
        )
    })
});
