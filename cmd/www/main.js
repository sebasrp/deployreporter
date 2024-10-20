import './style.css'
import 'vis-timeline/dist/vis-timeline-graph2d.min.css';
import $ from 'jquery';
import { DataSet } from 'vis-data';
import { Timeline} from 'vis-timeline';

var data;

function createColumnFilterDropdownList(listClass, classCheckbox, classTD, listItems){
  var listStr = '';
  listItems.sort().forEach(item=> {
    listStr += `
      <label class="label cursor-pointer">
        <input type="checkbox" checked="checked" class="checkbox `+classCheckbox+`" value="`+ item + `"/>
        <span class="label-text" align="left">` + item + `</span>
      </label>
`
  });
  document.getElementById(listClass).innerHTML = listStr;
  filterColumn(classCheckbox, classTD);
}

function filterColumn(checkboxClass, tdClass){
  $("."+checkboxClass).each(function() {$(this).on('change',function(){
    var $this = $(this);
    $("tbody > tr > td."+tdClass).each(function() {
        if ($this.val()==this.innerText){
          if($this.is(':checked')){
            $(this).parent().show();
          }else{
            $(this).parent().hide();
          }
        }
  });
})});
}

const getdata = async () => {
  const endpoint = "http://localhost:8080/deployments?limit=500",
        response = await fetch(endpoint),
        data = await response.json();

  var svcList = new Map();
  var environmentList = new Map();
  var countryList = new Map();
  var sourceList = new Map();

  var timelineItems = [];

  data.forEach(entry => {
      let {ID, Start, End, Operator, Service, Environment, Country, Source} = entry;
      var start= new Date(Start).toISOString();
      var end= new Date(End).toISOString();
      svcList.set(Service, "");
      environmentList.set(Environment, "");
      countryList.set(Country, "");
      sourceList.set(Source, "");

      tbody.innerHTML += `<tr>
          <td class="d-id">${ID}</td>
          <td class="d-start">${start}</td>
          <td class="d-end">${end}</td>
          <td class="d-operator">${Operator}</td>
          <td class="d-svc">${Service}</td>
          <td class="d-env">${Environment}</td>
          <td class="d-country">${Country}</td>
          <td class="d-source">${Source}</td>
      </tr>`;

      timelineItems.push({id:ID, content:Service, start:Start})
  });


  createColumnFilterDropdownList('svc-list', 'svc-checkbox', 'd-svc',  Array.from(svcList.keys()));
  createColumnFilterDropdownList('env-list', 'env-checkbox', 'd-env',  Array.from(environmentList.keys()));
  createColumnFilterDropdownList('country-list', 'country-checkbox', 'd-country',  Array.from(countryList.keys()));
  createColumnFilterDropdownList('source-list', 'source-checkbox', 'd-source',  Array.from(sourceList.keys()));

  // Create a Timeline
  var container = document.getElementById('timeline');
  var timelineOptions = {height:"200px"};
  var timeline = new Timeline(container, new DataSet(timelineItems), timelineOptions);
};

document.querySelector('#app').innerHTML = `
  <div>
    <!-- header -->
    <div class="navbar bg-base-100">
    <div class="flex-1">
      <a class="btn btn-ghost text-xl">Deployment Tracker</a>
    </div>
    <div class="flex-none">
      <label class="swap swap-rotate">
        <!-- this hidden checkbox controls the state -->
        <input type="checkbox" class="theme-controller" value="light" />

        <!-- sun icon -->
        <svg
          class="swap-on h-10 w-10 fill-current"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24">
          <path
            d="M5.64,17l-.71.71a1,1,0,0,0,0,1.41,1,1,0,0,0,1.41,0l.71-.71A1,1,0,0,0,5.64,17ZM5,12a1,1,0,0,0-1-1H3a1,1,0,0,0,0,2H4A1,1,0,0,0,5,12Zm7-7a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4A1,1,0,0,0,12,5ZM5.64,7.05a1,1,0,0,0,.7.29,1,1,0,0,0,.71-.29,1,1,0,0,0,0-1.41l-.71-.71A1,1,0,0,0,4.93,6.34Zm12,.29a1,1,0,0,0,.7-.29l.71-.71a1,1,0,1,0-1.41-1.41L17,5.64a1,1,0,0,0,0,1.41A1,1,0,0,0,17.66,7.34ZM21,11H20a1,1,0,0,0,0,2h1a1,1,0,0,0,0-2Zm-9,8a1,1,0,0,0-1,1v1a1,1,0,0,0,2,0V20A1,1,0,0,0,12,19ZM18.36,17A1,1,0,0,0,17,18.36l.71.71a1,1,0,0,0,1.41,0,1,1,0,0,0,0-1.41ZM12,6.5A5.5,5.5,0,1,0,17.5,12,5.51,5.51,0,0,0,12,6.5Zm0,9A3.5,3.5,0,1,1,15.5,12,3.5,3.5,0,0,1,12,15.5Z" />
        </svg>

        <!-- moon icon -->
        <svg
          class="swap-off h-10 w-10 fill-current"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24">
          <path
            d="M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z" />
        </svg>
      </label>
    </div>
  </div>
<!-- timeline -->
  <div id="timeline"></div>
     <!-- filters -->
    <div>
      <!-- table -->
      <table id="deployment-table" class="table table-zebra">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Start</th>
            <th scope="col">End</th>
            <th scope="col">Operator</th>
            <th scope="col">
              <details class="dropdown">
                <summary class="btn">Service</summary>
                <ul tabindex="0" id="svc-list" class="dropdown-content bg-base-100">
                </ul>
              </div>
            </th>
            <th scope="col">
              <details class="dropdown">
                <summary class="btn">Env</summary>
                <ul tabindex="0" id="env-list" class="dropdown-content bg-base-100">
                </ul>
              </div>
            </th>
            <th scope="col">
              <details class="dropdown">
                <summary class="btn">Country</summary>
                <ul tabindex="0" id="country-list" class="dropdown-content bg-base-100" >
                </ul>
              </div>
            </th>
            <th scope="col">
              <details class="dropdown">
                <summary class="btn">Source</summary>
                <ul tabindex="0" id="source-list" class="dropdown-content bg-base-100">
                </ul>
              </div>
            </th>
          </tr>
        </thead>
        <tbody id="tbody">
        </tbody>
      </table>
    </div>
  </div>
`

data=getdata();