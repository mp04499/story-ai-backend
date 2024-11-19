import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { StoryController } from './controllers';
import { StoryService } from './services/story.service';

@Module({
  imports: [ConfigModule.forRoot({ envFilePath: '.env.local' })],
  controllers: [StoryController],
  providers: [StoryService],
})
export class AppModule {}
