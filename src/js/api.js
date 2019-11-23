(function() {
  var DOM = {
    apiResponse: document.getElementById("api-response"),
    errorResponse: document.getElementById("error-response"),
  };

  fetch("/api/v1")
    .then(function(res) {
      if (!res.ok) {
        DOM.errorResponse.innerText = `${res.status}: ${res.statusText}`;
      } else {
        return res.json();
      }
    })
    .then(function(json) {
      DOM.apiResponse.innerText = JSON.stringify(json, null, 2);
    });
})();
