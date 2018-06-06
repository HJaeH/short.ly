



function UrlCtrl($scope, $http) {

  $scope.urls = [];
  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
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

         $http.post('/create', data, config)
            .success(function (data, status, headers, config) {
//                 refresh().then(function() {
//                        $scope.rawURL = '';
//                 })
            })
            .error(function (data, status, header, config) {
            });
//            $http.post('/create', {url: temp}).
//                    error(logError).
//                    success(function() {
//                      refresh().then(function() {
//                        $scope.rawURL = '';
//                      })
//                    });
         } else {
            // todo : display invalid url alert in view
            alert("유효한 URL이 아닙니다.");
         }
    };

}