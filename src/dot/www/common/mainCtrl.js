var D = angular.module('dot',[ngResource]);
D.controller('mainCtrl',['$scope'],function($scope,$timeout){
    $scope.hi = 'hello World';
});