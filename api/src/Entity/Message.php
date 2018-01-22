<?php
declare(strict_types=1);

namespace App\Entity;

use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Serializer\Annotation\Groups;

/**
 * @ORM\Entity
 */
class Message
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     *
     * @Groups({"conversation", "message"})
     */
    private $id;

    /**
     * @ORM\ManyToOne(targetEntity="User")
     * @ORM\JoinColumn(name="user_id", referencedColumnName="id")
     *
     * @Groups({"conversation", "message"})
     */
    private $from;

    /**
     * @ORM\Column(type="string")
     *
     * @Groups({"conversation", "message"})
     */
    private $body;

    /**
     * @ORM\ManyToOne(targetEntity="Conversation", inversedBy="messages")
     * @ORM\JoinColumn(name="conversation_id", referencedColumnName="id")
     *
     * @Groups({"message"})
     */
    private $conversation;

    /**
     * @Groups({"message"})
     */
    private $to;

    /**
     * @return mixed
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * @return User
     */
    public function getFrom(): User
    {
        return $this->from;
    }

    /**
     * @param User $from
     */
    public function setFrom(User $from): void
    {
        $this->from = $from;
    }

    /**
     * @return string
     */
    public function getBody():? string
    {
        return $this->body;
    }

    /**
     * @param string $body
     */
    public function setBody(string $body): void
    {
        $this->body = $body;
    }

    /**
     * @return string
     */
    public function getTo()
    {
        return $this->to;
    }

    /**
     * @param mixed $to
     */
    public function setTo($to): void
    {
        $this->to = $to;
    }

    /**
     * @return Conversation
     */
    public function getConversation()
    {
        return $this->conversation;
    }

    /**
     * @param Conversation $conversation
     */
    public function setConversation(Conversation $conversation)
    {
        $conversation->addMessage($this);
        $this->conversation = $conversation;
    }

    /**
     * @return array
     */
    public function denormalize(): array
    {
        return [
            'id' => $this->id,
            'from' => $this->getFrom()->getUsername(),
            'to' => $this->getTo(),
            'conversation' => $this->getConversation()->getId(),
            'body' => $this->getBody(),
        ];
    }
}
