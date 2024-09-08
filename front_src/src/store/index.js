// store/index.js
import { createStore } from 'vuex';

export default createStore({
  state: {
    loginData: {
      account: '',
      password: ''
    }
  },
  mutations: {
    setAccount(state, account) {
      state.loginData.account = account;
      console.log(state.loginData.account);
    },
    setPassword(state, password) {
      state.loginData.password = password;
      console.log(state.loginData.password);
    }
  },
  actions: {
    updateAccount({ commit }, account) {
      commit('setAccount', account);
    },
    updatePassword({ commit }, password) {
      commit('setPassword', password);
    }
  },
  getters: {
    getLoginData(state) {
      return state.loginData;
    },
    getAccount(state) {
        return state.loginData.account;
    }
  }
});
