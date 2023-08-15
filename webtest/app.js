fetch('https://')
.then(function (response){
    return response.json();
})
.then(function (data){
    appendData(data);
})
.catch(function (err){
    console.log('error: ' + err);
});
function appendData(data){
    var mainContainer = document.getElementById("myData");
    for (var i = 0; i < data.length; i++){
        var div = document.createElement("div");
        div.innerHTML = 'CourseID: ' + data[i].ID + ' ' + data[i].Name + ' ' + data[i].Price + ' ' + data[i].Instructor; 
        mainContainer.appendChild(div)
    }
}