//wywolanie na poczatku 
// if(localStorage.getItem("currentPage"))
//     getFilterCars(localStorage.getItem("currentPage"));
// else
//     getFilterCars();

function getFilterCars(currentPage=0){

    var filters,maxCarsPage=30;
    
    if(localStorage.getItem("filters"))
        filters=localStorage.getItem("filters");
    else
        filters="NULL";
    
    
    var cars=filterCars(filters);
    
   
    //var cars = [filters, '1', '3', '7' ];
    
    let temp, item, a, i;

    temp = document.getElementsByTagName("template")[0];
    item = temp.content.querySelector("div");
    if(cars!=null){
        // to do maksymalna liczba samochopdów dla kategorii
    for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {
        //console.log(i);
    if(i<Object.keys(cars).length){
      a = document.importNode(item, true);
      let elem=a.querySelectorAll("h3");
      elem[0].textContent=[cars[i].brand,cars[i].model].join(' ');
      let p=a.querySelectorAll("h5");
      p[0].textContent=["Dzienny koszt:",cars[i].dailyCost].join(' ');
      p[1].textContent=["Spalanie:",cars[i].fuelConsumption].join(' ');
      document.getElementById("cardGroup").appendChild(a);
    }
    }}
    else{
        let a=document.createElement("h4");//
        a.style.color='white';
        a.textContent="Brak samochodów dla wybranej kategorii.";
        
        document.getElementById("cardGroup").appendChild(a);
    }
    createFilterOptions();
    }


function filterChange(){
    document.getElementById("cardGroup").remove();
    localStorage.setItem("filters","filters");
    document.location.href = "user-rent.html";
}

function filterCars(filters){
    var jsonObj=JSON.parse(localStorage.getItem("allCars"));
    if(jsonObj)
        console.log(jsonObj[0].id);
    return jsonObj;
    //const map = new Map(Object.entries(JSON.parse(json)));
}

async function createFilterOptions(){
    
    const filterMap=new Map(JSON.parse(localStorage.getItem("allFilters")));
    createFOption(filterMap,"activeBrand","brand","brandList");
}

function createFOption(filterMap,buttonName,name,listName){
    let item=document.getElementById(buttonName);
    
    if(localStorage.getItem(name))
        item.innerText=localStorage.getItem(name)
    else
        item.innerText="Wszystkie";
    //item=document.getElementById("filterList");

    let i=0;
    for(i;i<filterMap.get(name).length;i++){
        a=document.createElement("li");
        a.appendChild(document.createTextNode(filterMap.get(name)[i]));
        a.classList.add("dropdown-item");
        document.getElementById(listName).appendChild(a);
    }
}