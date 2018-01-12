<?php
declare(strict_types=1);

namespace App\Event;

use App\Document\Conversation;
use App\Document\Message;
use Symfony\Component\EventDispatcher\Event;

class ConversationEvent extends Event
{
    const MESSAGE_CREATED = 'conversation.message.created';

    /**
     * @var Conversation
     */
    private $conversation;

    /**
     * @var Message
     */
    private $message;

    /**
     * ConversationEvent constructor.
     * @param Conversation $conversation
     * @param Message $message
     */
    public function __construct(Conversation $conversation, Message $message)
    {
        $this->conversation = $conversation;
        $this->message = $message;
    }

    /**
     * @return Conversation
     */
    public function getConversation(): Conversation
    {
        return $this->conversation;
    }

    /**
     * @return Message
     */
    public function getMessage(): Message
    {
        return $this->message;
    }
}
