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
        
        // if (now == tests.length-1) {
        //     var btm = test.children[4]
        //     btm.setAttribute("class", "submit")
        //     btm.textContent="Отправить"
        // }
        console.log(test, now)
        $('meta[name=now]').attr('content', parseInt(now)+1);
        }
    }
    )
    $('.submit').click(function (){
        //добавим в jQuery нужный метод
    $.fn.serializeObject = function()
    {
    var o = {};
    var a = this.serializeArray();
    $.each(a, function() {
        if (o[this.name] !== undefined) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
    };
    //используем новый метод на нужной форме
    var _object = $("#form_test").serializeObject();
    console.log(_object)
    })
});