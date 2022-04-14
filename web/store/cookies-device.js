export default {
  nuxtServerInit ({ state, commit }, { app, route, req, redirect }) {
    /**
     *  lấy cookies data và khở tạo lại state
     */
    const user = app.$cookies.get('user', { parseJSON: true })
    if (user) {
      commit('mutate', {
        property: 'currentUser',
        with: user
      })
    } else {
      commit('mutate', {
        property: 'currentUser',
        with: null
      })
    }

    initDevice({ commit }, { app, route, req, redirect })
  }
}

const initDevice = ({ commit }, { app, route }) => {
  // get device from cookies
  const device = app.$cookies.get('setDevice')
  if (device) {
    commit('mutate', {
      property: 'device',
      with: device
    })
  }
  if (route.query.device) {
    // set fake device cookies
    app.$cookies.set('setDevice', route.query.device, {
      path: '/',
      maxAge: 60 * 60 * 24 * 10
    })
    app.$cookies.set('device', route.query.device, {
      path: '/'
    })
    commit('mutate', {
      property: 'device',
      with: route.query.device
    })
    if (route.query.device === 'mobile') {
      app.$cookies.set('os', 'ios', {
        path: '/'
      })
    } else {
      app.$cookies.set('os', '', {
        path: '/'
      })
    }
  }
  // set device cookie by real device
  if (app.$device.isDesktop) {
    app.$cookies.set('device', 'desktop', {
      path: '/'
    })
    if (!device) {
      commit('mutate', {
        property: 'device',
        with: 'desktop'
      })
    }
    app.$cookies.set('os', 'desktop', { path: '/' })
  } else {
    app.$cookies.set('device', 'mobile', {
      path: '/'
    })
    if (app.$device.isIos) {
      app.$cookies.set('os', 'ios', {
        path: '/'
      })
    } else {
      app.$cookies.set('os', 'android', {
        path: '/'
      })
    }
    if (!device) {
      commit('mutate', {
        property: 'device',
        with: 'mobile'
      })
    }
  }
}
