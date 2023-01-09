if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1) == 'user-rent.html'||window.location.href.substring(window.location.href.lastIndexOf('/') + 1) == 'employee-rent.html'||window.location.href.substring(window.location.href.lastIndexOf('/') + 1)=='index=rent.html') {
    var currentLoc=(window.location.href.substring(window.location.href.lastIndexOf('/') + 1));
document.getElementById("startDate").addEventListener("change", function() {
    var input = this.value;
    localStorage.setItem("startDate",input);
    document.location.href = currentLoc;
});
document.getElementById("endDate").addEventListener("change", function() {
    var input = this.value;
    localStorage.setItem("endDate",input);
    document.location.href = currentLoc;
});
}
function padTo2Digits(num) {
    return num.toString().padStart(2, '0');
}

function formatDate(date) {
    return [
      date.getFullYear(),
      padTo2Digits(date.getMonth() + 1),
      padTo2Digits(date.getDate()),
    ].join('-');
  }

async function getFilterCars(currentPage=0){
    var start=new Date(JSON.stringify(localStorage.getItem("startDate")));
    var end=new Date(JSON.stringify(localStorage.getItem("endDate")));
    start.setDate(start.getDate()+1);
    end.setDate(end.getDate()+1);
    var tempDate;
    if(!isNaN(start)){
        document.getElementById("startDate").valueAsDate=start;
    }
    else{
        tempDate=new Date();
        localStorage.setItem("startDate",formatDate(tempDate));
        document.getElementById("startDate").valueAsDate = tempDate;
    }
    if(!isNaN(end)){
        document.getElementById("endDate").valueAsDate = end;
    }
    else{
        tempDate=new Date();
        localStorage.setItem("endDate",formatDate(tempDate));
        document.getElementById("endDate").valueAsDate = tempDate;
    }

    var filters;
    if(localStorage.getItem("filters"))
        filters=JSON.parse(localStorage.getItem("filters"));
    else{
        // active filters
        filters={"brand":"Wszystkie","model":"Wszystkie","type":"Wszystkie","color":"Wszystkie"};
        localStorage.setItem("filters",JSON.stringify(filters));
    }

    const carsDate = {
        start_date: localStorage.getItem("startDate"),
        end_date: localStorage.getItem("endDate")
    }

    var filteredCars;

    Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/vehicles/available","POST",carsDate)).then(cars=>{
        filteredCars=filterCars(cars,filters);
        console.log(filteredCars)
        makeFilters(cars);
        createFilterOptions();
        printFilteredCars(filteredCars,currentPage);
    }).catch( err => {
    console.log('error: '+ err);
    alert("Złe daty!");
  });
}

function printFilteredCars(filteredCars,currentPage){
    let temp, item, a, i,maxCarsPage=30;

    temp = document.getElementsByTagName("template")[0];
    item = temp.content.querySelector("div");

    if(Object.keys(filteredCars).length!=0){
    for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {

    if(i<Object.keys(filteredCars).length){
        
        a = document.importNode(item, true);
        let elem=a.querySelectorAll("h3");
        elem[0].textContent=[filteredCars[i].brand,filteredCars[i].model].join(' ');
        let p=a.querySelectorAll("h5");
        p[0].textContent=["Dzienny koszt:",filteredCars[i].dailyCost].join(' ');
        p[1].textContent=["Spalanie:",filteredCars[i].fuelConsumption].join(' ');

        a.querySelector("img").src=getPhoto(filteredCars[i].brand,filteredCars[i].model);
        if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)!=="index-rent.html"){
            let b=a.querySelectorAll("button");
            b[0].id=filteredCars[i].id;
            b[0].addEventListener('click', function handleClick(event) {
                localStorage.setItem("currentCar",this.id);
                var currentLoc=(window.location.href.substring(window.location.href.lastIndexOf('/') + 1))
                if(currentLoc=="employee-rent.html")
                    document.location.href = "employee-checkout.html";
                else    
                    document.location.href="user-checkout.html";
                rentCar();
    });
}
        document.getElementById("cardGroup").appendChild(a);
    }
    if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)==="employee-rent.html"){
        let b=a.querySelectorAll("button");
        b[1].id=b[0].id;
    }
    }}
    else{
        let a=document.createElement("h4");
        a.style.color='white';
        a.textContent="Brak samochodów dla wybranej kategorii.";
        
        document.getElementById("cardGroup").appendChild(a);
    }
}


// zwraca mape ktora zawiera wszystkie wystapienia searchValue
function getByValue(map, searchValue) {
    const final = new Map();
    for (let [key, value] of map.entries()) {
        //key = pojedynczy samochod
        Object.entries(value).forEach(([k,v]) => {
        // gdy jedna wartosc rowna
        if(searchValue==v){
            final.set(key,value);
        }
    });
    }
    return final;
  }

// filtrowanie samochodów
function filterCars(data,filters){
    let map= new Map(Object.entries(data));
   
   let filteredCars=new Map(map);
   
   Object.entries(filters).forEach(([key,value]) => {
    if(value!="Wszystkie"){
        filteredCars=getByValue(filteredCars,value);
    }
   });
    
   const res=Object.fromEntries(filteredCars);

   let returnArray=[];
   Object.entries(res).forEach(([key,value])=>{
    returnArray.push(value);
   });   
    return returnArray;
}
function changeFilter(name){
    // przekazywana wartosc w name to "kategoria kliknietaOpcja"
    const words = name.split(' ');
    var filters=JSON.parse(localStorage.getItem("filters"));
    filters[words[0]]=words[1];
    localStorage.setItem("filters",JSON.stringify(filters));
    var currentLoc=("currentPageLocation",window.location.href.substring(window.location.href.lastIndexOf('/') + 1))
    document.location.href = currentLoc;
}
async function createFilterOptions(){
    const filterMap=new Map(JSON.parse(localStorage.getItem("allFilters")));
    let filters=JSON.parse(localStorage.getItem("filters"));
    createFOption(filterMap,"activeBrand","brand",filters.brand,"brandList");
    createFOption(filterMap,"activeModel","model",filters.model,"modelList");
    createFOption(filterMap,"activeType","type",filters.type,"typeList");
    createFOption(filterMap,"activeColor","color",filters.color,"colorList");
}

// generowanie opcji do wybrania w filtrowaniu
function createFOption(filterMap,buttonName,fName,name,listName){
    let item=document.getElementById(buttonName);
    if(name)
        item.innerText=name;
    else
        item.innerText="Wszystkie";
    let i=0,temp;
    for(i;i<filterMap.get(fName).length;i++){
        a=document.createElement("li");
        a.appendChild(document.createTextNode(filterMap.get(fName)[i]));
        a.id=fName+" "+filterMap.get(fName)[i];

        if(filterMap.get(fName)[i]==name){
            a.classList.add("disabled");
        }
        a.addEventListener('click', function handleClick(event) {
                changeFilter(this.id);
        });
        a.classList.add("dropdown-item");
        document.getElementById(listName).appendChild(a);
    }
}

function makeFilters(data){
    const brand = ["Wszystkie",...new Set(data.map(item => item.brand))];
    const model = ["Wszystkie",...new Set(data.map(item => item.model))];
    const type = ["Wszystkie",...new Set(data.map(item => item.type))]; 
    const color = ["Wszystkie",...new Set(data.map(item => item.color))];
    const map1= new Map();
    map1.set('brand',brand);
    map1.set('model',model);
    map1.set('type',type);
    map1.set('color',color);
    localStorage.setItem('allFilters',JSON.stringify(Array.from(map1.entries())));
}

  


