import type { Component } from 'solid-js';
import { createSignal, For } from 'solid-js';

import { Memo } from '../components/Memo';

export const Home: Component = () => {
  const [memos, setMemos] = createSignal<any[]>([]);

  (async () => {
    const res = await fetch('/api/memos');
    const memos = await res.json();

    setMemos(memos);
  })();

  return (
    <For each={memos()}>{memo => <Memo content={memo.content} createdAt={memo.createdAt} />}</For>
  )
}
