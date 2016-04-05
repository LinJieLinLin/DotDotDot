var M = angular.module('myApp', []);
var directive = M;
M.controller('myCtrl', ['$scope', function($scope) {}]);
M.directive('my2048', function() {
    return {
        restrict: 'E',
        replace: true,
        transclude: true,
        templateUrl: '2048.html',
        scope: {
            a: '=',
            b: '='
        },
        controller: function($scope, $element, $rootElement) {
            $scope.home = [];
            for (var i = 0; i < 4; i++) {
                $scope.home[i] = [];
                for (var j = 0; j < 4; j++) {
                    $scope.home[i][j] = j;
                }
            }
            $rootElement.find("body")[0].style.backgroundColor = "#EDDFF3";
            console.log(
                $rootElement.find("body")[0].style.backgroundColor);
            console.log($element);
        }
    };
});
