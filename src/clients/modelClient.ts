import axios from 'axios';
import { Message, ModelResponse } from 'src/types/modelClientTypes';

export class ModelClient {
  private readonly instance = axios.create({
    baseURL: process.env.modelUrl,
    timeout: 10000,
  });

  async completion(messages: Message[]): Promise<ModelResponse> {
    try {
      const response: ModelResponse = await this.instance.post(
        '/v1/chat/completions',
        {
          model: 'GawdSB/story_model',
          messages,
          max_new_tokens: 2000,
          top_k: 50,
          top_p: 0.95,
          num_return_sequences: 1,
          do_sample: true,
        },
      );

      return response;
    } catch (error) {
      throw Error('Error in calling LLM Model. Message: ' + error);
    }
  }
}
