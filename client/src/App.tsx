import type { Component } from 'solid-js';
import { Routes, Route } from "solid-app-router";

import { Home } from './pages/Home';

const App: Component = () => {
  return (
    <>
      <div class="uk-navbar-container tm-navbar-container uk-sticky uk-sticky-fixed uk-active uk-sticky-below">
        <div class="uk-container uk-container-expand">
          <nav class="uk-navbar">
            <div class="uk-navbar-left">
              <h1 class="uk-navbar-item uk-logo">
                Generate ORiginal Memo
              </h1>
            </div>
            <div class="uk-navbar-right">
              <button class="uk-button uk-button-primary uk-margin-small-right" uk-toggle="target: #message" type="button"> New Memo </button>
            </div>
          </nav>
        </div>
      </div>
      <div id="message" uk-modal>
        <div class="uk-modal-dialog uk-modal-body">
          <h2 class="uk-modal-title">New Memo</h2>
          <form>
            <textarea class="uk-textarea"></textarea>
          </form>
          <button class="uk-button uk-button-primary uk-margin-small-right" type="button">Create</button>
          <button class="uk-modal-close-outside" type="button" uk-close />
        </div>
      </div>
      <Routes>
        <Route path="/" component={Home} />
      </Routes>
    </>
  );
};

export default App;
