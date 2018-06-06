



function UrlCtrl($scope, $http, $location) {

  $scope.urls = [];
  $scope.localURLs = [];
  $scope.tasks = [];
  var logError = function(data, status) {
    console.log('code '+status+': '+data);

  };
   $scope.records = [
          "Alfreds Futterkiste",
          "Berglunds snabbköp",
          "Centro comercial Moctezuma",
          "Ernst Handel",
      ]

   var refresh = function() {
      return $http.get('/create/').
        success(function(data) { $scope.localURLs = data; }).
        error(logError);
    };


  $scope.addURL = function() {
        // check url validation.
         if(/^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|www\.)[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/.test($scope.rawURL)){
             var data = $.param({
                             url: $scope.rawURL,
                         });
             var config = {
                 headers : {
                     'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8;'
                 }
             }

             $http.post('/create', data, config).
             then(function(response) {
                $scope.status = response.status;
                $scope.data = response.data;
                $scope.localURLs.push($scope.data);
                $scope.records.push($scope.data);
                refresh().then(function() {

                        })

             }, function(response) {
                $scope.data = response.data || 'Request failed';
                $scope.status = response.status;
           });

         } else {
            alert("유효한 URL이 아닙니다.");
         }
    };

}