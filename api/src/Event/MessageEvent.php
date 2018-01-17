<?php
declare(strict_types=1);

namespace App\Event;

use App\Entity\Message;
use Symfony\Component\EventDispatcher\Event;

class MessageEvent extends Event
{
    const CREATED = 'message.created';

    /**
     * @var Message
     */
    private $message;

    /**
     * @param Message $message
     */
    public function __construct(Message $message)
    {
        $this->message = $message;
    }

    /**
     * @return Message
     */
    public function getMessage(): Message
    {
        return $this->message;
    }
}
