D.controller('mainCtrl', ['$scope', '$timeout', 'S', function($scope, $timeout, S) {
    $scope.D = {};
    $scope.req = {};
    $scope.month = (new Date).getMonth() + 1;
    $scope.buttonName = '下单';
    $scope.allPrice = 0;
    var lastMenu = angular.fromJson(localStorage.lastMenu);
    if (!lastMenu) {
        lastMenu = {};
        lastMenu.id = 0;
        lastMenu.name = '';
    }
    $scope.myMenu = {
        id: lastMenu.id || 0,
        name: lastMenu.name || '',
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
    //下单列表 
    $scope.req.getList = function(params, filter) {
        var option = {
            method: 'GET',
            url: '/d/list',
            params: params
        };
        return S.R(angular.extend(option, filter || {}));
    };
    //下单 
    $scope.req.setList = function(params, filter) {
        var option = {
            method: 'POST',
            url: '/d/list',
            params: params
        };
        return S.R(angular.extend(option, filter || {}));
    };
    $scope.setList = function() {
        if ($scope.buttonName === '已下单') {
            // return;
        }
        if ($scope.myMenu.id < 1 || !angular.isArray($scope.myMenu.menuList)) {
            return;
        }
        var len = $scope.myMenu.menuList.length;
        var data = [];
        for (var i = 0; i < len; i++) {
            data[i] = {};
            data[i].uid = $scope.myMenu.id;
            data[i].mid = $scope.myMenu.menuList[i].Id;
        }
        $scope.req.setList({ data: JSON.stringify(data) }).then(function(rs) {
            if (rs.code === 0) {
                $scope.buttonName = '已下单';
            }
        }, function(e) {
            console.log(e);
            console.log(e.data.msg);
        });
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
    $scope.today = new Date().toLocaleDateString().replace(/\//g, '-');
    // $scope.today = '2016-02-25';
    $scope.req.getList({ time: $scope.today }).then(function(rs) {
        console.log(rs);
        $scope.D.buyList = rs.data;
        angular.forEach(rs.data.Menu, function(v) {
            $scope.allPrice += +v.Price;
        });
    }, function(e) {
        console.log(e);
    });
    //选中数据时添加到已选列表
    $scope.select = function(argData, argType) {
        if (argType === 'user') {
            var lastUser = JSON.parse(localStorage.lastMenu);
            if (lastUser.id === argData.Id) {
                $scope.myMenu.id = argData.Id;
                $scope.myMenu.name = argData.Name;
            } else {
                $scope.myMenu = {
                    id: argData.Id,
                    name: argData.Name,
                    menuList: [],
                    price: 0
                };
            }
        } else if (argType === 'menu') {
            $scope.myMenu.menuList.push(argData);
            $scope.myMenu.price += argData.Price;
        }
        if ($scope.myMenu.price > 20) {
            // $scope.myMenu.priceColor.color = 'red';
        }
        localStorage.lastMenu = JSON.stringify($scope.myMenu);
    };
    $scope.delMenu = function(argIndex) {
        $scope.myMenu.price = $scope.myMenu.price - $scope.myMenu.menuList[argIndex].Price;
        $scope.myMenu.menuList.splice(argIndex, 1);
        $scope.myMenu.price <= 0 ? 0 : $scope.myMenu.price;
    };
}]);
