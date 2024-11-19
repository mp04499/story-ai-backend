export interface ModelResponse {
  object: string;
  id: string;
  created: Date;
  model: string;
  system_fingerprint: string;
  choices: Choice[];
  usage: Usage;
}

export interface Choice {
  index: number;
  message: Message;
  finish_reason: string;
}

export interface Message {
  role: string;
  content: string;
}

export interface Usage {
  prompt_tokens: number;
  completion_tokens: number;
  total_tokens: number;
}

export interface ModelRequest {
  model: 'GawdSB/story_model';
  messages: Message[];
  max_new_token: number;
  top_k: number;
  top_p: number;
  do_sample: boolean;
}
