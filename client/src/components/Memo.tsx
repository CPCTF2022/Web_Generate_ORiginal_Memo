import type { Component } from 'solid-js';

export const Memo: Component<any> = (props) => {
  return (
    <a class="uk-card">
      <div class="uk-card-body">{props.content}</div>
      <div class="uk-card-footer">
        <p class="uk-text-meta uk-margin-remove-top"><time datetime={props.createdAt}>{props.createdAt}</time></p>
      </div>
    </a>
  )
}
