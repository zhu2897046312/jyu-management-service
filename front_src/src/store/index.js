// store/index.js
import { createStore } from 'vuex';

export default createStore({
  state: {
    loginData: JSON.parse(localStorage.getItem('loginData')) || {  // 从 localStorage 中恢复数据
      account: '',
      password: ''
    }
  },
  mutations: {
    setAccount(state, account) {
      state.loginData.account = account;
      localStorage.setItem('loginData', JSON.stringify(state.loginData));  // 每次修改时更新 localStorage
    },
    setPassword(state, password) {
      state.loginData.password = password;
      localStorage.setItem('loginData', JSON.stringify(state.loginData));  // 每次修改时更新 localStorage
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
