
function test(filters="NULL",maxCars=15,currentPage=1){
    var myArr = ['Ram', 'Shyam', 'Sita', 'Gita' ];
    
    var temp, item, a, i;
    temp = document.getElementsByTagName("template")[0];

    item = temp.content.querySelector("div");
    for (i = 0; i < myArr.length; i++) {
      a = document.importNode(item, true);
      a.getElementById("name").textContent=myArr[i];
      //a.textContent += myArr[i];
      document.getElementById("cardGroup").appendChild(a);
      //document.body.appendChild(a);
    }
}