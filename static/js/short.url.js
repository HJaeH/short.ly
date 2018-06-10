



function UrlCtrl($scope, $http, $window) {

  $scope.urls = [];
  $scope.localURLs = [];
  $scope.tasks = [];
  var logError = function(data, status) {
    console.log('code '+status+': '+data);

  };

   var refresh = function() {
//      return $http.get('/list').
//        success(function(data) { $scope.localURLs = data; }).
//        error(logError);
    };


   $scope.redir = function(data){
       $window.location.replace(data);
     }

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
                var found = false;
                for(var i = 0; i < $scope.localURLs.length; i++) {
                    if ($scope.localURLs[i].OriginalURL == $scope.data.OriginalURL ) {
                        if ( $scope.localURLs[i].OriginalURL != undefined){
                            found = true;
                            break;
                        }

                    }
                }
                if (!found){
                    var temp = $scope.data;
                    $scope.localURLs.push(temp);
                } else {
                    console.log($scope.data, "1111")
                }

             }, function(response) {
                $scope.data = response.data || 'Request failed';
                $scope.status = response.status;
           });

         } else {
            alert("유효한 URL이 아닙니다.");
         }
    };

}