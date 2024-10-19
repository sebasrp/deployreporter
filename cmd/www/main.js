import './style.css'
import { setupCounter } from './counter.js'


const getdata = async () => {
  const endpoint = "http://localhost:8080/deployments?limit=500",
        response = await fetch(endpoint),
        data = await response.json();

 data.forEach(entry => {
    let {ID, Start, End, Service, Environment, Country, Source} = entry;
    var start= new Date(Start).toISOString();
    var end= new Date(End).toISOString();

    tbody.innerHTML += `<tr>
        <td>${ID}</td>
        <td>${start}</td>
        <td>${end}</td>
        <td>${Service}</td>
        <td>${Environment}</td>
        <td>${Country}</td>
        <td>${Source}</td>
    </tr>`;
 });
}


document.querySelector('#app').innerHTML = `
  <div>
    <h1>Deployment tracker</h1>
    <div class="overflow-x-auto">
      <table class="table table-zebra">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Start</th>
            <th scope="col">End</th>
            <th scope="col">Service</th>
            <th scope="col">Environment</th>
            <th scope="col">Country</th>
            <th scope="col">Source</th>
          </tr>
        </thead>
        <tbody id="tbody">
        </tbody>
      </table>
    </div>
  </div>
`

getdata();