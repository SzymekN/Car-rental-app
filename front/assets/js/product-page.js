//to do:  getAllFilters has to be in filter change, delete local variables when car is rented (start and end date)
if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1) == 'user-rent.html') {
document.getElementById("startDate").addEventListener("change", function() {
    var input = this.value;
    //console.log(input);
    localStorage.setItem("startDate",input);
    document.location.href = "user-rent.html";
});
document.getElementById("endDate").addEventListener("change", function() {
    var input = this.value;
    //console.log(input);
    localStorage.setItem("endDate",input);
    document.location.href = "user-rent.html";
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
    //alert("F")
    //console.log("Page location is " + window.location.href.substring(window.location.href.lastIndexOf('/') + 1))
    var start=new Date(JSON.stringify(localStorage.getItem("startDate")));
    var end=new Date(JSON.stringify(localStorage.getItem("endDate")));
    start.setDate(start.getDate()+1);
    end.setDate(end.getDate()+1);

    var tempDate;
    if(!isNaN(start)){
        document.getElementById("startDate").valueAsDate=start;
    }
    else{
        tempDate=new Date()
        localStorage.setItem("startDate",formatDate(tempDate))
        document.getElementById("startDate").valueAsDate = tempDate
    }
    if(!isNaN(end)){
        document.getElementById("endDate").valueAsDate = end;
    }
    else{
        tempDate=new Date()
        localStorage.setItem("endDate",formatDate(tempDate))
        document.getElementById("endDate").valueAsDate = tempDate
    }

    var filters;
    if(localStorage.getItem("filters"))
        filters=JSON.parse(localStorage.getItem("filters"));
    else{
        // active filters
        filters={"brand":"Wszystkie","model":"Wszystkie","type":"Wszystkie","color":"Wszystkie"};
        localStorage.setItem("filters",JSON.stringify(filters));
    }

    var filteredCars;
    Promise.resolve(getAvailableCars()).then(cars=>{
        //localStorage.setItem("allCars",JSON.stringify(cars))
        filteredCars=filterCars(cars,filters);
        makeFilters(cars);
        createFilterOptions();
        printFilteredCars(filteredCars,currentPage);
});
}

function printFilteredCars(filteredCars,currentPage){
    let temp, item, a, i,maxCarsPage=30;

    temp = document.getElementsByTagName("template")[0];
    item = temp.content.querySelector("div");

    if(Object.keys(filteredCars).length!=0){
        // to do maksymalna liczba samochopdów dla kategorii
    for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {

    //if(i<Object.keys(cars).length){
    if(i<Object.keys(filteredCars).length){
        
        a = document.importNode(item, true);
        let elem=a.querySelectorAll("h3");
        elem[0].textContent=[filteredCars[i].brand,filteredCars[i].model].join(' ');
        let p=a.querySelectorAll("h5");
        p[0].textContent=["Dzienny koszt:",filteredCars[i].dailyCost].join(' ');
        p[1].textContent=["Spalanie:",filteredCars[i].fuelConsumption].join(' ');
        let b=a.querySelectorAll("button");
        b[0].id=filteredCars[i].id;
        b[0].addEventListener('click', function handleClick(event) {
            localStorage.setItem("currentCar",this.id);
            //console.log(this.id);
            document.location.href = "car-rent.html";
            rentCar();
    });
        document.getElementById("cardGroup").appendChild(a);
    }
    }}
    else{
        let a=document.createElement("h4");//
        a.style.color='white';
        a.textContent="Brak samochodów dla wybranej kategorii.";
        
        document.getElementById("cardGroup").appendChild(a);
    }
}
// function filterChange(){
//     document.getElementById("cardGroup").remove();
//     localStorage.setItem("filters","filters");
//     document.location.href = "user-rent.html";
// }

// zwraca mape ktora zawiera wszystkie wystapienia searchValue
function getByValue(map, searchValue) {
    const final = new Map();
    for (let [key, value] of map.entries()) {
        //key = pojedynczy samochod
        Object.entries(value).forEach(([k,v]) => {
        //console.log(v);
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
    //const jsonObj=JSON.parse(localStorage.getItem("allCars"));
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
    //console.log(returnArray);
    
    return returnArray;
}
function changeFilter(name){
    // przekazywana wartosc w name to "kategoria kliknietaOpcja"
    const words = name.split(' ');
    var filters=JSON.parse(localStorage.getItem("filters"));
    filters[words[0]]=words[1];
    localStorage.setItem("filters",JSON.stringify(filters));
    document.location.href = "user-rent.html";
}
async function createFilterOptions(){
    
    const filterMap=new Map(JSON.parse(localStorage.getItem("allFilters")));
    //console.log("createFilteroptions"+Object.entries(filterMap))
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
    //console.log(filterMap)
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
    //const data = new Map(Object.entries(JSON.parse(jsonData)));
    
    const brand = ["Wszystkie",...new Set(data.map(item => item.brand))];
    const model = ["Wszystkie",...new Set(data.map(item => item.model))];
    const type = ["Wszystkie",...new Set(data.map(item => item.type))]; 
    const color = ["Wszystkie",...new Set(data.map(item => item.color))];
    const map1= new Map();
    map1.set('brand',brand);
    map1.set('model',model);
    map1.set('type',type);
    map1.set('color',color);
    //console.log(map1.get('color'));
    localStorage.setItem('allFilters',JSON.stringify(Array.from(map1.entries())));
}

  
async function getAvailableCars() {
    const carsDate = {
        start_date: localStorage.getItem("startDate"),
        end_date: localStorage.getItem("endDate")
    }
    //dzialalo przed
    var target="http://192.168.33.50:8200/api/v1/vehicles/available";
    return new Promise((res, rej) => {                       // return a promise
    fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(carsDate),
      headers: {
        "Content-Type": "application/json; charset=UTF-8",
        "Content-Length":"217",
        "Authorization":"Bearer "+localStorage.getItem("token")
      }}).then((r) => {   // fetch the resourse
        // const isJson = r.headers.get('content-type')?.includes('application/json')
        const data =r.json();
        if(!r.ok)
        {
          const error = (data && data.message) || r.status;
          return Promise.reject(error);
        }
          //res(r);                                      // resolve promise if success
          return res(data);
      }).then(res.toString).catch( err => {
          return rej(err);                         // don't try again 
      });                                              // again until no more tries
  });
  
  getData.then(data=>{
    localStorage.setItem("allCars",JSON.stringify(data))
    //makeFilters(data);
    console.log("F"+Object.values(data));
  }).catch(err=>console.log(err));
}


