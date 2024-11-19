import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { StoryController } from './controllers';

@Module({
  imports: [ConfigModule.forRoot({ envFilePath: '.env.local' })],
  controllers: [AppController, StoryController],
  providers: [AppService],
})
export class AppModule {}
