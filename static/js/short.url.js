/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function TaskCtrl($scope, $http) {
  $scope.tasks = [];

  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/task/').
      success(function(data) { $scope.tasks = data.Tasks; }).
      error(logError);
  };

  $scope.addTodo = function() {
    $scope.working = true;
    $http.post('/task/', {Title: $scope.todoText}).
      error(logError).
      success(function() {
        refresh().then(function() {

          $scope.working = false;
          $scope.todoText = '';
        })
      });
  };

  $scope.toggleDone = function(task) {
    data = {ID: task.ID, Title: task.Title, Done: !task.Done}
    $http.put('/task/'+task.ID, data).
      error(logError).
      success(function() { task.Done = !task.Done });
  };

  $scope.shortenURL = function(originalURL) {
      data = {ID: task.ID, Title: task.Title, Done: !task.Done}
      $http.post('/shorten/'+task.ID, data).
        error(logError).
        success(function() { task.Done = !task.Done });
    };
  $scope.addURL = function() {
      $scope.working = true;
      $http.post('/shorten/', {Title: $scope.originalURL}).
        error(logError).
        success(function() {
          refresh().then(function() {
            $scope.working = false;
            $scope.todoText = '';
          })
        });
    };


  refresh().then(function() { $scope.working = false; });
}


function UrlCtrl($scope, $http) {

  $scope.urls = [];
  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/task/').
      success(function(data) { $scope.tasks = data.Tasks; }).
      error(logError);
  };

  $scope.shortenURL = function(originalURL) {
      data = {ID: task.ID, Title: task.Title, Done: !task.Done}
      $http.post('/shorten/'+task.ID, data).
        error(logError).
        success(function() { task.Done = !task.Done });
    };
  $scope.addURL = function() {
      $scope.working = true;

      var rawURL = $scope.originalURL;
        // check url validation.
         if(/^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|www\.)[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/.test(rawURL)){
            $http.post('/shorten/', {Title: $scope.originalURL}).
                    error(logError).
                    success(function() {
                      refresh().then(function() {
                        $scope.todoText = '';
                      })
                    });
         } else {
            // todo : display invalid url alert in view
            alert("유효한 URL이 아닙니다.");
         }



    };


  refresh().then(function() { $scope.working = false; });
}