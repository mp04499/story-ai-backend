import { Body, Controller, Get, Post } from '@nestjs/common';
import { StoryService } from 'src/services/story.service';
import { StoryResponse } from 'src/types/storyTypes';

@Controller('story')
export class StoryController {
  service = new StoryService();

  @Get()
  getTest(): string {
    return 'test';
  }

  @Post()
  async getStory(@Body('text') text: string): Promise<StoryResponse> {
    this.service.text = text;
    const messages = await this.service.callModel();

    return {
      data: messages,
    };
  }
}
