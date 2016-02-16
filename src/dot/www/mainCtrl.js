D.controller('mainCtrl', ['$scope', '$timeout', 'S', function($scope, $timeout, S) {
    $scope.D = {};
    $scope.req = {};
    $scope.month = (new Date).getMonth() + 1;
    $scope.myMenu = {
        id: '',
        name: '',
        menuList: [],
        price: 0,
        priceColor: { color: 'yellow' }
    };
    //用户
    $scope.req.getUser = function(params, filter) {
        var option = {
            method: 'GET',
            url: '/d/name',
            params: params
        };
        return S.R(angular.extend(option, filter || {}));
    };
    //菜单 
    $scope.req.getMenu = function(params, filter) {
        var option = {
            method: 'GET',
            url: '/d/menu',
            params: params
        };
        return S.R(angular.extend(option, filter || {}));
    };

    $scope.req.getUser().then(function(rs) {
        $scope.D.user = rs.data;
    }, function(e) {
        console.log(e);
    });
    $scope.req.getMenu().then(function(rs) {
        $scope.D.menu = rs.data;
        $scope.D.menuDefault = [];
        $scope.D.menuMonth = [];
        $scope.D.menuLitle = [];
        angular.forEach(rs.data, function(d, index) {
            console.log(d);
            if (d.Type === 0) {
                $scope.D.menuDefault.push(d);
            } else if (d.Type === $scope.month) {
                $scope.D.menuMonth.push(d);
            } else if (d.Type < 0) {
                $scope.D.menuLitle.push(d);
            }
        });
    }, function(e) {
        console.log(e);
    });
    //选中数据时添加到已选列表
    $scope.select = function(argData, argType) {
        if (argType === 'user') {
            $scope.myMenu.id = argData.Id;
            $scope.myMenu.name = argData.Name;
        } else if (argType === 'menu') {
            $scope.myMenu.menuList.push(argData);
            $scope.myMenu.price += argData.Price;
        }
        if ($scope.myMenu.price > 20) {
            $scope.myMenu.priceColor.color = 'red';
        }
    };
}]);
