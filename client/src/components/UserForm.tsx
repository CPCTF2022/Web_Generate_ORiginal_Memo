import type { Component } from 'solid-js';

export const UserForm: Component<any> = (props) => {
  return (
    <form>
      <fieldset class="uk-fieldset">
        <div class="uk-margin">
          <input class="uk-input" type="text" placeholder="Name" />
          <input class="uk-input" type="password" placeholder="Password" />
        </div>
      </fieldset>
      <button class="uk-button uk-button-primary uk-margin-small-right" type="button" onClick={props.handler}>{props.label}</button>
    </form>
  )
}
