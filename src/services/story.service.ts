import { Injectable } from '@nestjs/common';
import type { Message, ResponseData } from '../types/modelClientTypes';
import { ModelClient } from 'src/clients/modelClient';

@Injectable()
export class StoryService {
  _text: string;
  _history: string[] = [];
  client = new ModelClient();

  set text(value: string) {
    if (this._text && this._text.length) {
      this._history.push(this._text);
    }
    this._text = value;
  }

  get text() {
    return this._text;
  }

  async callModel(): Promise<ResponseData> {
    const messages = this.formatText();
    const res = await this.client.completion(messages);
    return res;
  }

  formatText(): Message[] {
    return [
      {
        role: 'system',
        content: 'You are a imaginative story telling chat bot',
      },
      { role: 'user', content: this._text },
    ];
  }
}
